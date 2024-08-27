import { productAction, productActionTypes } from "./actions";
import { Product } from "./product";


export interface DispatcherProps {
    addProduct: (payload: Product) => void
    editProduct: (payload: Product) => void
    removeProduct: (id: number) => void
}

const productStoreDispatcher = {
    addProduct: (payload: Product) => ({
        type: productActionTypes.add, payload: payload
    }),
    editProduct: (payload: Product) => ({
        type: productActionTypes.edit, payload: payload
    }),
    removeProduct: (id: number) => ({
        type: productActionTypes.delete, payload: {
            id: id
        }
    })
}

export default productStoreDispatcher