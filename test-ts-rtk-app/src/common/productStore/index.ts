import { DispatcherProps } from "./dispatcher"
import { Product } from "./product"


export type ProductStoreProps = StateProps & DispatcherProps

export interface StateProps {
    elements: Product[]
}

export const productStoreState = (state: any) => ({
    elements: state.products
})