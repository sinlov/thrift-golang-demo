namespace java com.sinlov.demo.rpc

// Rpc Service
service RpcService {
    // RPC call
    list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
}