import grpc
from vfl.v1 import coordinator_pb2 as pb
from vfl.v1 import coordinator_pb2_grpc as stubs

def main():
    # insecure channel for dev; switch to TLS later
    channel = grpc.insecure_channel("localhost:8443")
    client = stubs.CoordinatorStub(channel)
    resp = client.RegisterParty(pb.RegisterPartyRequest(party_id="partyA"))
    print("session_id:", resp.session_id)
    print("echo:", resp.echo)

if __name__ == "__main__":
    main()
