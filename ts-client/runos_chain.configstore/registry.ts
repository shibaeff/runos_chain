import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSetPort } from "./types/runoschain/configstore/tx";
import { MsgCreateHostsDatabase } from "./types/runoschain/configstore/tx";
import { MsgUpdateHostsDatabase } from "./types/runoschain/configstore/tx";
import { MsgDeleteHostsDatabase } from "./types/runoschain/configstore/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/runos_chain.configstore.MsgSetPort", MsgSetPort],
    ["/runos_chain.configstore.MsgCreateHostsDatabase", MsgCreateHostsDatabase],
    ["/runos_chain.configstore.MsgUpdateHostsDatabase", MsgUpdateHostsDatabase],
    ["/runos_chain.configstore.MsgDeleteHostsDatabase", MsgDeleteHostsDatabase],
    
];

export { msgTypes }