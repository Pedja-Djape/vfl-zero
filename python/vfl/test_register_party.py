from pathlib import Path
import grpc
from vfl.v1 import coordinator_pb2 as pb
from vfl.v1 import coordinator_pb2_grpc as stubs

def get_channel_creds():
    base_path = Path(__file__).resolve().parents[2] / "certs"
    root_certs = (base_path / "dev-ca.crt").read_bytes()
    client_cert = (base_path / "client.crt").read_bytes()
    client_key = (base_path / "client.key").read_bytes()

    creds = grpc.ssl_channel_credentials(
        root_certificates=root_certs,
        private_key=client_key,
        certificate_chain=client_cert,
    )
    return creds

def main():
    creds = get_channel_creds()
    channel = grpc.secure_channel("localhost:8443", creds)
    client = stubs.CoordinatorStub(channel)
    resp = client.RegisterParty(pb.RegisterPartyRequest(party_id="partyA"))
    print("session_id:", resp.session_id)
    print("echo:", resp.echo)

if __name__ == "__main__":
    main()
