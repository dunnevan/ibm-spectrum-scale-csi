package v1

import "strings"

const (
	ScaleCore        = "scale-core"
	ScaleGUI         = "scale-gui"
	ScalePMCollector = "scale-pmcollector"

	IBMSpectrumScaleCR          = "ibm-spectrum-scale"
	IBMSpectrumScaleCore        = "ibm-spectrum-scale-core"
	IBMSpectrumScaleGUI         = "ibm-spectrum-scale-gui"
	IBMSpectrumScalePMCollector = "ibm-spectrum-scale-pmcollector"
	IBMSpectrumScaleCABundle    = "ibm-spectrum-scale-cabundle"

	ScaleInternalStorageClass = "ibm-spectrum-scale-internal"

	VolumeMountCABundle = "/etc/ssl/service"
)

var (
	ScaleGUISecretPostfix = "-gui-" + strings.ToLower(ContainerOperator)
)
