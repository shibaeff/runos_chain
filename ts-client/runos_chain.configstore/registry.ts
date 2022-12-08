import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSetPort } from "./types/runoschain/configstore/tx";
import { MsgGetPort } from "./types/runoschain/configstore/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/runos_chain.configstore.MsgSetPort", MsgSetPort],
    ["/runos_chain.configstore.MsgGetPort", MsgGetPort],
    
];

export { msgTypes }