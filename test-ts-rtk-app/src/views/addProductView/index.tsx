import { connect, useDispatch, useSelector } from "react-redux"
import ProductForm from "../../components/productForm/ProductForm"
import { Product } from "../../common/_productStore/product"
import { useAppDispatch } from "../.."
import { addProduct } from "../../common/productStore/slice"


const AddProductView = () => {
    
    const {products} = useSelector((state: any) => state)
    const dispatch = useAppDispatch();

    const handleFormSubmit = (data: Product) => {

        const lastElementId = products[products.length - 1]?.id || 0

        dispatch(addProduct({
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