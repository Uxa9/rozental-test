import { Key, useEffect, useState } from "react"
import ElementView from "./elementView"
import { connect, shallowEqual, useDispatch, useSelector } from "react-redux"
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore"
import Table from "../../components/table/Table"
import { Product } from "../../common/productStore/product"
import productStoreReducer from "../../common/productStore/reducer"


const TableView: React.FC = () => {
    const productList = useSelector((state: any) => state.products)    

    return (
        <>
            <Table elements={productList.map((element: Product, index: Key) => {
                return (
                    <ElementView
                        key={index}
                        product={element}
                    />
                )
            })} />
        </>
    )
}

export default TableView