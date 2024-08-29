

export enum cartActionTypes {
    add = "ADD_ITEM",
    remove = "REMOVE_ITEM",
    wipe = "WIPE_CART"
}

export interface cartAction {
    type: cartActionTypes,
    payload?: { id: number }
}