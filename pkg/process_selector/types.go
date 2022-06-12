package processselector

type LoadT int64

const (
	UndefinedLoad LoadT = iota
	Loaded
	NotFound
	BadSetting
	Error
	Masked
)

func (a LoadT) String() string {
	switch a {
	case UndefinedLoad:
		return "undefined"
	case Loaded:
		return "loaded"
	case NotFound:
		return "not found"
	case BadSetting:
		return "bad setting"
	case Error:
		return "error"
	case Masked:
		return "masked"
	}
	return "undefined"
}

type ActiveT int64

const (
	UndefinedActive ActiveT = iota
	Active
	Reloading
	Inactive
	Failed
	Activating
	Deactivating
)

func (l ActiveT) String() string {
	switch l {
	case UndefinedActive:
		return "undefined"
	case Active:
		return "active"
	case Reloading:
		return "reloading"
	case Inactive:
		return "inactive"
	case Failed:
		return "failed"
	case Activating:
		return "activating"
	case Deactivating:
		return "deactivating"
	}
	return "undefined"
}

type Unit struct {
	Name        string
	Load        LoadT
	Active      ActiveT
	Sub         string
	Description string
}
