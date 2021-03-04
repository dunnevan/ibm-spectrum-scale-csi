package v1

const (
	//Data Key for GuiUserSecret.  Value is comma separated list of user groups.
	Group = "group"

	Administrator     = "Administrator"
	CsiAdmin          = "CsiAdmin"
	ContainerOperator = "ContainerOperator" //old versions of GUI used: "CnssOperator" which we never supported for release
)
