import { Client, registry, MissingWalletError } from 'runos_chain-client-ts'

import { HostsDatabase } from "runos_chain-client-ts/runos_chain.configstore/types"
import { Params } from "runos_chain-client-ts/runos_chain.configstore/types"


export { HostsDatabase, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				HostsDatabase: {},
				HostsDatabaseAll: {},
				GetPort: {},
				
				_Structure: {
						HostsDatabase: getStructure(HostsDatabase.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getHostsDatabase: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.HostsDatabase[JSON.stringify(params)] ?? {}
		},
				getHostsDatabaseAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.HostsDatabaseAll[JSON.stringify(params)] ?? {}
		},
				getGetPort: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetPort[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: runos_chain.configstore initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.RunosChainConfigstore.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryHostsDatabase({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.RunosChainConfigstore.query.queryHostsDatabase( key.dpid,  key.mac)).data
				
					
				commit('QUERY', { query: 'HostsDatabase', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryHostsDatabase', payload: { options: { all }, params: {...key},query }})
				return getters['getHostsDatabase']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryHostsDatabase API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryHostsDatabaseAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.RunosChainConfigstore.query.queryHostsDatabaseAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.RunosChainConfigstore.query.queryHostsDatabaseAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'HostsDatabaseAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryHostsDatabaseAll', payload: { options: { all }, params: {...key},query }})
				return getters['getHostsDatabaseAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryHostsDatabaseAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetPort({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.RunosChainConfigstore.query.queryGetPort( key.dpid,  key.mac)).data
				
					
				commit('QUERY', { query: 'GetPort', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetPort', payload: { options: { all }, params: {...key},query }})
				return getters['getGetPort']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetPort API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateHostsDatabase({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.RunosChainConfigstore.tx.sendMsgCreateHostsDatabase({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateHostsDatabase:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateHostsDatabase:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateHostsDatabase({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.RunosChainConfigstore.tx.sendMsgUpdateHostsDatabase({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateHostsDatabase:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateHostsDatabase:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteHostsDatabase({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.RunosChainConfigstore.tx.sendMsgDeleteHostsDatabase({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteHostsDatabase:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteHostsDatabase:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetPort({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.RunosChainConfigstore.tx.sendMsgSetPort({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetPort:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetPort:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateHostsDatabase({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.RunosChainConfigstore.tx.msgCreateHostsDatabase({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateHostsDatabase:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateHostsDatabase:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateHostsDatabase({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.RunosChainConfigstore.tx.msgUpdateHostsDatabase({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateHostsDatabase:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateHostsDatabase:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteHostsDatabase({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.RunosChainConfigstore.tx.msgDeleteHostsDatabase({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteHostsDatabase:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteHostsDatabase:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetPort({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.RunosChainConfigstore.tx.msgSetPort({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetPort:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetPort:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
