package mock_nodeattestor

//go:generate sh -c "$GOPATH/bin/mockgen github.com/spiffe/spire/proto/agent/nodeattestor NodeAttestor,NodeAttestorServer,NodeAttestor_FetchAttestationDataClient > nodeattestor_mock.go"
