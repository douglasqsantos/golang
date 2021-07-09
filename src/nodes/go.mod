module nodes

go 1.16

replace dqs.io/utils/auth => ../dqs.io/utils/auth

replace dqs.io/utils/errors => ../dqs.io/utils/errors

replace dqs.io/utils/msgs => ../dqs.io/utils/msgs

require (
	dqs.io/utils/auth v0.0.0-00010101000000-000000000000 // indirect
	dqs.io/utils/errors v0.0.0-00010101000000-000000000000 // indirect
	dqs.io/utils/msgs v0.0.0-00010101000000-000000000000 // indirect
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)
