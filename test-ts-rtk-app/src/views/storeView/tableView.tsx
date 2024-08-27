import { useEffect, useState } from "react"
import ElementView from "./elementView"
import { connect, shallowEqual, useSelector } from "react-redux"
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore"
import Table from "../../components/table/Table"
import { Product } from "../../common/productStore/product"


const TableView: React.FC<ProductStoreProps> = (props) => {
    
    const [elements, setElements] = useState<Product[]>([])
    
    useEffect(() => {
        setElements(props.elements)
    }, [props.elements])

    return (
        <>
            <Table elements={elements.map((element, index) => {
                return (
                    <ElementView
                        key={index}
                        own={element}
                    />
                )
            })} />
        </>
    )
}

export default connect<StateProps, DispatcherProps>(
    productStoreState,
    productStoreDispatcher,
)(TableView)