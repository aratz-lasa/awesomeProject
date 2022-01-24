using Go = import "/go.capnp";
@0xef83879a531f9bf3;
$Go.package("main");
$Go.import("github.com/aratz-lasa/awesomeProject/mudp");


struct MudpRequest {
    peer @0 :Data;
    namespace @1 :Text;
}

struct MudpResponse {
    namespace @0 :Text;
    envelope @1 :List(Data);
}

struct MudpPacket {
    union {
        request @0 :MudpRequest;
        response @1 :MudpResponse;
    }  
}