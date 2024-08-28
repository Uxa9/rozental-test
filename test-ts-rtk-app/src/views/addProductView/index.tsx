import { connect, useDispatch, useSelector } from "react-redux"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore"
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher"
import Button from "../../components/button/Button"
import ProductForm from "../../components/productForm/ProductForm"
import { Product } from "../../common/productStore/product"

type Props = {
    product?: Product
}

const AddProductView: React.FC<Props> = (props) => {
    
    const productList = useSelector((state: any) => state.products)
    const dispatch = useDispatch()

    const handleFormSubmit = (data: Product) => {

        const lastElementId = productList[productList.length - 1]?.id || 0

        dispatch(productStoreDispatcher.addProduct({
            id: lastElementId + 1,
            name: data.name,
            price: data.price,
        }))
        
    }

    return (
        <ProductForm resultHandler={handleFormSubmit} />
    )
}

export default AddProductView