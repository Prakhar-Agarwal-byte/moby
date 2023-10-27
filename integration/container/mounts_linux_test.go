package container // import "github.com/Prakhar-Agarwal-byte/moby/integration/container"

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	containertypes "github.com/Prakhar-Agarwal-byte/moby/api/types/container"
	mounttypes "github.com/Prakhar-Agarwal-byte/moby/api/types/mount"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/network"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/versions"
	"github.com/Prakhar-Agarwal-byte/moby/client"
	"github.com/Prakhar-Agarwal-byte/moby/integration/internal/container"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/parsers/kernel"
	"github.com/Prakhar-Agarwal-byte/moby/testutil"
	"github.com/moby/sys/mount"
	"github.com/moby/sys/mountinfo"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/fs"
	"gotest.tools/v3/poll"
	"gotest.tools/v3/skip"
)

func TestContainerNetworkMountsNoChown(t *testing.T) {
	// chown only applies to Linux bind mounted volumes; must be same host to verify
	skip.If(t, testEnv.IsRemoteDaemon)

	ctx := setupTest(t)

	tmpDir := fs.NewDir(t, "network-file-mounts", fs.WithMode(0o755), fs.WithFile("nwfile", "network file bind mount", fs.WithMode(0o644)))
	defer tmpDir.Remove()

	tmpNWFileMount := tmpDir.Join("nwfile")

	config := containertypes.Config{
		Image: "busybox",
	}
	hostConfig := containertypes.HostConfig{
		Mounts: []mounttypes.Mount{
			{
				Type:   "bind",
				Source: tmpNWFileMount,
				Target: "/etc/resolv.conf",
			},
			{
				Type:   "bind",
				Source: tmpNWFileMount,
				Target: "/etc/hostname",
			},
			{
				Type:   "bind",
				Source: tmpNWFileMount,
				Target: "/etc/hosts",
			},
		},
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	assert.NilError(t, err)
	defer cli.Close()

	ctrCreate, err := cli.ContainerCreate(ctx, &config, &hostConfig, &network.NetworkingConfig{}, nil, "")
	assert.NilError(t, err)
	// container will exit immediately because of no tty, but we only need the start sequence to test the condition
	err = cli.ContainerStart(ctx, ctrCreate.ID, containertypes.StartOptions{})
	assert.NilError(t, err)

	// Check that host-located bind mount network file did not change ownership when the container was started
	// Note: If the user specifies a mountpath from the host, we should not be
	// attempting to chown files outside the daemon's metadata directory
	// (represented by `daemon.repository` at init time).
	// This forces users who want to use user namespaces to handle the
	// ownership needs of any external files mounted as network files
	// (/etc/resolv.conf, /etc/hosts, /etc/hostname) separately from the
	// daemon. In all other volume/bind mount situations we have taken this
	// same line--we don't chown host file content.
	// See GitHub PR 34224 for details.
	info, err := os.Stat(tmpNWFileMount)
	assert.NilError(t, err)
	fi := info.Sys().(*syscall.Stat_t)
	assert.Check(t, is.Equal(fi.Uid, uint32(0)), "bind mounted network file should not change ownership from root")
}

func TestMountDaemonRoot(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon)

	ctx := setupTest(t)
	apiClient := testEnv.APIClient()
	info, err := apiClient.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range []struct {
		desc        string
		propagation mounttypes.Propagation
		expected    mounttypes.Propagation
	}{
		{
			desc:        "default",
			propagation: "",
			expected:    mounttypes.PropagationRSlave,
		},
		{
			desc:        "private",
			propagation: mounttypes.PropagationPrivate,
		},
		{
			desc:        "rprivate",
			propagation: mounttypes.PropagationRPrivate,
		},
		{
			desc:        "slave",
			propagation: mounttypes.PropagationSlave,
		},
		{
			desc:        "rslave",
			propagation: mounttypes.PropagationRSlave,
			expected:    mounttypes.PropagationRSlave,
		},
		{
			desc:        "shared",
			propagation: mounttypes.PropagationShared,
		},
		{
			desc:        "rshared",
			propagation: mounttypes.PropagationRShared,
			expected:    mounttypes.PropagationRShared,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			test := test
			t.Parallel()

			ctx := testutil.StartSpan(ctx, t)

			propagationSpec := fmt.Sprintf(":%s", test.propagation)
			if test.propagation == "" {
				propagationSpec = ""
			}
			bindSpecRoot := info.DockerRootDir + ":" + "/foo" + propagationSpec
			bindSpecSub := filepath.Join(info.DockerRootDir, "containers") + ":/foo" + propagationSpec

			for name, hc := range map[string]*containertypes.HostConfig{
				"bind root":    {Binds: []string{bindSpecRoot}},
				"bind subpath": {Binds: []string{bindSpecSub}},
				"mount root": {
					Mounts: []mounttypes.Mount{
						{
							Type:        mounttypes.TypeBind,
							Source:      info.DockerRootDir,
							Target:      "/foo",
							BindOptions: &mounttypes.BindOptions{Propagation: test.propagation},
						},
					},
				},
				"mount subpath": {
					Mounts: []mounttypes.Mount{
						{
							Type:        mounttypes.TypeBind,
							Source:      filepath.Join(info.DockerRootDir, "containers"),
							Target:      "/foo",
							BindOptions: &mounttypes.BindOptions{Propagation: test.propagation},
						},
					},
				},
			} {
				t.Run(name, func(t *testing.T) {
					hc := hc
					t.Parallel()

					ctx := testutil.StartSpan(ctx, t)

					c, err := apiClient.ContainerCreate(ctx, &containertypes.Config{
						Image: "busybox",
						Cmd:   []string{"true"},
					}, hc, nil, nil, "")
					if err != nil {
						if test.expected != "" {
							t.Fatal(err)
						}
						// expected an error, so this is ok and should not continue
						return
					}
					if test.expected == "" {
						t.Fatal("expected create to fail")
					}

					defer func() {
						if err := apiClient.ContainerRemove(ctx, c.ID, containertypes.RemoveOptions{Force: true}); err != nil {
							panic(err)
						}
					}()

					inspect, err := apiClient.ContainerInspect(ctx, c.ID)
					if err != nil {
						t.Fatal(err)
					}
					if len(inspect.Mounts) != 1 {
						t.Fatalf("unexpected number of mounts: %+v", inspect.Mounts)
					}

					m := inspect.Mounts[0]
					if m.Propagation != test.expected {
						t.Fatalf("got unexpected propagation mode, expected %q, got: %v", test.expected, m.Propagation)
					}
				})
			}
		})
	}
}

func TestContainerBindMountNonRecursive(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon)
	skip.If(t, versions.LessThan(testEnv.DaemonAPIVersion(), "1.40"), "BindOptions.NonRecursive requires API v1.40")
	skip.If(t, testEnv.IsRootless, "cannot be tested because RootlessKit executes the daemon in private mount namespace (https://github.com/rootless-containers/rootlesskit/issues/97)")

	ctx := setupTest(t)

	tmpDir1 := fs.NewDir(t, "tmpdir1", fs.WithMode(0o755),
		fs.WithDir("mnt", fs.WithMode(0o755)))
	defer tmpDir1.Remove()
	tmpDir1Mnt := filepath.Join(tmpDir1.Path(), "mnt")
	tmpDir2 := fs.NewDir(t, "tmpdir2", fs.WithMode(0o755),
		fs.WithFile("file", "should not be visible when NonRecursive", fs.WithMode(0o644)))
	defer tmpDir2.Remove()

	err := mount.Mount(tmpDir2.Path(), tmpDir1Mnt, "none", "bind,ro")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mount.Unmount(tmpDir1Mnt); err != nil {
			t.Fatal(err)
		}
	}()

	// implicit is recursive (NonRecursive: false)
	implicit := mounttypes.Mount{
		Type:     "bind",
		Source:   tmpDir1.Path(),
		Target:   "/foo",
		ReadOnly: true,
	}
	recursive := implicit
	recursive.BindOptions = &mounttypes.BindOptions{
		NonRecursive: false,
	}
	recursiveVerifier := []string{"test", "-f", "/foo/mnt/file"}
	nonRecursive := implicit
	nonRecursive.BindOptions = &mounttypes.BindOptions{
		NonRecursive: true,
	}
	nonRecursiveVerifier := []string{"test", "!", "-f", "/foo/mnt/file"}

	apiClient := testEnv.APIClient()
	containers := []string{
		container.Run(ctx, t, apiClient, container.WithMount(implicit), container.WithCmd(recursiveVerifier...)),
		container.Run(ctx, t, apiClient, container.WithMount(recursive), container.WithCmd(recursiveVerifier...)),
		container.Run(ctx, t, apiClient, container.WithMount(nonRecursive), container.WithCmd(nonRecursiveVerifier...)),
	}

	for _, c := range containers {
		poll.WaitOn(t, container.IsSuccessful(ctx, apiClient, c), poll.WithDelay(100*time.Millisecond))
	}
}

func TestContainerVolumesMountedAsShared(t *testing.T) {
	// Volume propagation is linux only. Also it creates directories for
	// bind mounting, so needs to be same host.
	skip.If(t, testEnv.IsRemoteDaemon)
	skip.If(t, testEnv.IsUserNamespace)
	skip.If(t, testEnv.IsRootless, "cannot be tested because RootlessKit executes the daemon in private mount namespace (https://github.com/rootless-containers/rootlesskit/issues/97)")

	ctx := setupTest(t)

	// Prepare a source directory to bind mount
	tmpDir1 := fs.NewDir(t, "volume-source", fs.WithMode(0o755),
		fs.WithDir("mnt1", fs.WithMode(0o755)))
	defer tmpDir1.Remove()
	tmpDir1Mnt := filepath.Join(tmpDir1.Path(), "mnt1")

	// Convert this directory into a shared mount point so that we do
	// not rely on propagation properties of parent mount.
	if err := mount.MakePrivate(tmpDir1.Path()); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mount.Unmount(tmpDir1.Path()); err != nil {
			t.Fatal(err)
		}
	}()
	if err := mount.MakeShared(tmpDir1.Path()); err != nil {
		t.Fatal(err)
	}

	sharedMount := mounttypes.Mount{
		Type:   mounttypes.TypeBind,
		Source: tmpDir1.Path(),
		Target: "/volume-dest",
		BindOptions: &mounttypes.BindOptions{
			Propagation: mounttypes.PropagationShared,
		},
	}

	bindMountCmd := []string{"mount", "--bind", "/volume-dest/mnt1", "/volume-dest/mnt1"}

	apiClient := testEnv.APIClient()
	containerID := container.Run(ctx, t, apiClient, container.WithPrivileged(true), container.WithMount(sharedMount), container.WithCmd(bindMountCmd...))
	poll.WaitOn(t, container.IsSuccessful(ctx, apiClient, containerID), poll.WithDelay(100*time.Millisecond))

	// Make sure a bind mount under a shared volume propagated to host.
	if mounted, _ := mountinfo.Mounted(tmpDir1Mnt); !mounted {
		t.Fatalf("Bind mount under shared volume did not propagate to host")
	}

	mount.Unmount(tmpDir1Mnt)
}

func TestContainerVolumesMountedAsSlave(t *testing.T) {
	// Volume propagation is linux only. Also it creates directories for
	// bind mounting, so needs to be same host.
	skip.If(t, testEnv.IsRemoteDaemon)
	skip.If(t, testEnv.IsUserNamespace)
	skip.If(t, testEnv.IsRootless, "cannot be tested because RootlessKit executes the daemon in private mount namespace (https://github.com/rootless-containers/rootlesskit/issues/97)")

	ctx := testutil.StartSpan(baseContext, t)

	// Prepare a source directory to bind mount
	tmpDir1 := fs.NewDir(t, "volume-source", fs.WithMode(0o755),
		fs.WithDir("mnt1", fs.WithMode(0o755)))
	defer tmpDir1.Remove()
	tmpDir1Mnt := filepath.Join(tmpDir1.Path(), "mnt1")

	// Prepare a source directory with file in it. We will bind mount this
	// directory and see if file shows up.
	tmpDir2 := fs.NewDir(t, "volume-source2", fs.WithMode(0o755),
		fs.WithFile("slave-testfile", "Test", fs.WithMode(0o644)))
	defer tmpDir2.Remove()

	// Convert this directory into a shared mount point so that we do
	// not rely on propagation properties of parent mount.
	if err := mount.MakePrivate(tmpDir1.Path()); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mount.Unmount(tmpDir1.Path()); err != nil {
			t.Fatal(err)
		}
	}()
	if err := mount.MakeShared(tmpDir1.Path()); err != nil {
		t.Fatal(err)
	}

	slaveMount := mounttypes.Mount{
		Type:   mounttypes.TypeBind,
		Source: tmpDir1.Path(),
		Target: "/volume-dest",
		BindOptions: &mounttypes.BindOptions{
			Propagation: mounttypes.PropagationSlave,
		},
	}

	topCmd := []string{"top"}

	apiClient := testEnv.APIClient()
	containerID := container.Run(ctx, t, apiClient, container.WithTty(true), container.WithMount(slaveMount), container.WithCmd(topCmd...))

	// Bind mount tmpDir2/ onto tmpDir1/mnt1. If mount propagates inside
	// container then contents of tmpDir2/slave-testfile should become
	// visible at "/volume-dest/mnt1/slave-testfile"
	if err := mount.Mount(tmpDir2.Path(), tmpDir1Mnt, "none", "bind"); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mount.Unmount(tmpDir1Mnt); err != nil {
			t.Fatal(err)
		}
	}()

	mountCmd := []string{"cat", "/volume-dest/mnt1/slave-testfile"}

	if result, err := container.Exec(ctx, apiClient, containerID, mountCmd); err == nil {
		if result.Stdout() != "Test" {
			t.Fatalf("Bind mount under slave volume did not propagate to container")
		}
	} else {
		t.Fatal(err)
	}
}

// Regression test for #38995 and #43390.
func TestContainerCopyLeaksMounts(t *testing.T) {
	ctx := setupTest(t)

	bindMount := mounttypes.Mount{
		Type:   mounttypes.TypeBind,
		Source: "/var",
		Target: "/hostvar",
		BindOptions: &mounttypes.BindOptions{
			Propagation: mounttypes.PropagationRSlave,
		},
	}

	apiClient := testEnv.APIClient()
	cid := container.Run(ctx, t, apiClient, container.WithMount(bindMount), container.WithCmd("sleep", "120s"))

	getMounts := func() string {
		t.Helper()
		res, err := container.Exec(ctx, apiClient, cid, []string{"cat", "/proc/self/mountinfo"})
		assert.NilError(t, err)
		assert.Equal(t, res.ExitCode, 0)
		return res.Stdout()
	}

	mountsBefore := getMounts()

	_, _, err := apiClient.CopyFromContainer(ctx, cid, "/etc/passwd")
	assert.NilError(t, err)

	mountsAfter := getMounts()

	assert.Equal(t, mountsBefore, mountsAfter)
}

func TestContainerBindMountRecursivelyReadOnly(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon)
	skip.If(t, versions.LessThan(testEnv.DaemonAPIVersion(), "1.44"), "requires API v1.44")

	ctx := setupTest(t)

	// 0o777 for allowing rootless containers to write to this directory
	tmpDir1 := fs.NewDir(t, "tmpdir1", fs.WithMode(0o777),
		fs.WithDir("mnt", fs.WithMode(0o777)))
	defer tmpDir1.Remove()
	tmpDir1Mnt := filepath.Join(tmpDir1.Path(), "mnt")
	tmpDir2 := fs.NewDir(t, "tmpdir2", fs.WithMode(0o777),
		fs.WithFile("file", "should not be writable when recursively read only", fs.WithMode(0o666)))
	defer tmpDir2.Remove()

	if err := mount.Mount(tmpDir2.Path(), tmpDir1Mnt, "none", "bind"); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := mount.Unmount(tmpDir1Mnt); err != nil {
			t.Fatal(err)
		}
	}()

	rroSupported := kernel.CheckKernelVersion(5, 12, 0)

	nonRecursiveVerifier := []string{`/bin/sh`, `-xc`, `touch /foo/mnt/file; [ $? = 0 ]`}
	forceRecursiveVerifier := []string{`/bin/sh`, `-xc`, `touch /foo/mnt/file; [ $? != 0 ]`}

	// ro (recursive if kernel >= 5.12)
	ro := mounttypes.Mount{
		Type:     mounttypes.TypeBind,
		Source:   tmpDir1.Path(),
		Target:   "/foo",
		ReadOnly: true,
		BindOptions: &mounttypes.BindOptions{
			Propagation: mounttypes.PropagationRPrivate,
		},
	}
	roAsStr := ro.Source + ":" + ro.Target + ":ro,rprivate"
	roVerifier := nonRecursiveVerifier
	if rroSupported {
		roVerifier = forceRecursiveVerifier
	}

	// Non-recursive
	nonRecursive := ro
	nonRecursive.BindOptions = &mounttypes.BindOptions{
		ReadOnlyNonRecursive: true,
		Propagation:          mounttypes.PropagationRPrivate,
	}

	// Force recursive
	forceRecursive := ro
	forceRecursive.BindOptions = &mounttypes.BindOptions{
		ReadOnlyForceRecursive: true,
		Propagation:            mounttypes.PropagationRPrivate,
	}

	apiClient := testEnv.APIClient()

	containers := []string{
		container.Run(ctx, t, apiClient, container.WithMount(ro), container.WithCmd(roVerifier...)),
		container.Run(ctx, t, apiClient, container.WithBindRaw(roAsStr), container.WithCmd(roVerifier...)),

		container.Run(ctx, t, apiClient, container.WithMount(nonRecursive), container.WithCmd(nonRecursiveVerifier...)),
	}

	if rroSupported {
		containers = append(containers,
			container.Run(ctx, t, apiClient, container.WithMount(forceRecursive), container.WithCmd(forceRecursiveVerifier...)),
		)
	}

	for _, c := range containers {
		poll.WaitOn(t, container.IsSuccessful(ctx, apiClient, c), poll.WithDelay(100*time.Millisecond))
	}
}
