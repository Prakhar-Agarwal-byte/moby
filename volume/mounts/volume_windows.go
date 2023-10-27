package mounts // import "github.com/Prakhar-Agarwal-byte/moby/volume/mounts"

func (p *linuxParser) HasResource(m *MountPoint, absolutePath string) bool {
	return false
}
