package params

// SequoiaBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sequoia network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SequoiaBootnodes = []string{
	"enode://54533f5cf71e470f52e25841fe14af966f02b4da75846057791a83c2bf51791fd78e6a3ce77669faef595826f25219d180fd1a7cb124f9ed63d10104d8a8575d@bootsequoia.edevapps.com.br:30303",
	"enode://f9088ef47ebeb29150e9f8ba051b0e5e8428e32a65f6eafa9a4a590334139e3d7be75528bb2d943d278d8256fd8e0d1923299b655a65568fd271ab7bb55d49ae@187.89.87.254:30303",
}

// Once Sequoia network has DNS discovery set up,
// this value can be configured.
// var SequoiaDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
