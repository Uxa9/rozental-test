import { Product } from "./product";


export enum productActionTypes {
    add = "ADD_PRODUCT",
    edit = "EDIT_PRODUCT",
    delete = "DELETE_PRODUCT",
}

export interface productAction {
    type: productActionTypes,
    payload: Product
}