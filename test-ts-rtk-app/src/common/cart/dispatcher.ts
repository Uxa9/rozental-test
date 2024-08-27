import { cartActionTypes } from "./actions"

export interface DispatcherProps {
    addToCart: (id: number) => void
    removeFromCart: (id: number) => void
    wipeCart: () => void
}

const CartDispatcher = {
    addToCart: (id: number) => ({
        type: cartActionTypes.add, payload: {id: id}
    }),
    removeFromCart: (id: number) => ({
        type: cartActionTypes.remove, payload: {id: id}
    }),
    wipeCart: () => ({
        type: cartActionTypes.wipe
    })
}

export default CartDispatcher