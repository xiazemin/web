namespace go echo

struct Help{
1:optional string name;
2:required i32 time;
}
struct EchoReq {
    1: string msg;
    2: required string trace;
    3: optional string option;
    4: optional Help help;
}

struct EchoRes {
    1: string msg;
}

service Echo {
    EchoRes echo(1: EchoReq req);
}