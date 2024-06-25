package object

type FileStatus string

const (
	StatusNew      FileStatus = "new file:"
	StatusModified FileStatus = "modified:"
)

type FileInfo struct {
	ModifiedAt string
	Path       string
	Status     FileStatus
	Staging    bool
}

type FileInfoArr []FileInfo

func (a FileInfoArr) Len() int           { return len(a) }
func (a FileInfoArr) Less(i, j int) bool { return a[i].Path < a[j].Path }
func (a FileInfoArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func FromFileMetadataToFileInfo(f Metadata) FileInfo {
	var status FileStatus
	var staging bool

	if f.Repo == empty {
		status = StatusNew
	} else if f.Work != f.Repo {
		status = StatusModified
	}

	if f.Work == f.Stage {
		staging = true
	}

	info := FileInfo{
		ModifiedAt: f.Mtime,
		Path:       f.Path,
		Status:     status,
		Staging:    staging,
	}

	return info
}
