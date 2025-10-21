from pathlib import Path
import grpc
from vfl.v1 import coordinator_pb2 as pb
from vfl.v1 import coordinator_pb2_grpc as stubs

def main():
    ca_path = Path(__file__).resolve().parents[2] / "certs" / "dev-ca.crt"
    root_certs = ca_path.read_bytes()
    creds = grpc.ssl_channel_credentials(root_certificates=root_certs)

    channel = grpc.secure_channel("localhost:8443", creds)
    client = stubs.CoordinatorStub(channel)
    resp = client.RegisterParty(pb.RegisterPartyRequest(party_id="partyA"))
    print("session_id:", resp.session_id)
    print("echo:", resp.echo)

if __name__ == "__main__":
    main()
