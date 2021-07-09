module dqs.io/utils/auth

go 1.16

replace dqs.io/utils/errors => ../errors

replace dqs.io/utils/msgs => ../msgs

require (
	dqs.io/utils/errors v0.0.0-00010101000000-000000000000 // indirect
	dqs.io/utils/msgs v0.0.0-00010101000000-000000000000 // indirect
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)
