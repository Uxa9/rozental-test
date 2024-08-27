import { connect } from "react-redux"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore"
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher"
import Button from "../../components/button/Button"
import ProductForm from "../../components/productForm/ProductForm"
import { Product } from "../../common/productStore/product"


const AddProductView: React.FC<ProductStoreProps> = (props) => {

    const handleFormSubmit = (data: Product) => {

        const lastElementId = props.elements[props.elements.length - 1]?.id || 0

        props.addProduct({
            id: lastElementId + 1,
            name: data.name,
            price: data.price,
        })
        
    }

    return (
        <ProductForm resultHandler={handleFormSubmit} />
    )
}

export default connect<StateProps, DispatcherProps>(
    productStoreState,
    productStoreDispatcher,
)(AddProductView)