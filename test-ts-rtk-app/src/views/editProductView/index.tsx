import { useLocation, useNavigate } from "react-router-dom"
import { Product } from "../../common/productStore/product"
import ProductForm from "../../components/productForm/ProductForm"
import { useEffect, useState } from "react";
import { connect, useDispatch, useSelector } from "react-redux";
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore";
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher";


const EditProductView = () => {
    
    const location = useLocation();
    const navigate = useNavigate();

    const productList = useSelector((state: any) => state.products)
    const dispatch = useDispatch()

    const [currentElement, setCurrentElement] = useState<Product>({id: 0, name: "", price: 0})

    useEffect(() => {
        let curEl = productList.find((el: Product) => el.id === parseInt(location.pathname.split('/').pop() || ""))!
        !curEl && navigate("/");
        setCurrentElement(curEl);
    }, [])

    const handleFormSubmit = (data: Product) => {

        dispatch(productStoreDispatcher.editProduct({
            id: data.id,
            name: data.name,
            price: data.price,
        }))

        navigate(`/product/${data.id}`)
    }

    return (
        <ProductForm 
            product={currentElement}
            resultHandler={handleFormSubmit} 
        />
    )
}

export default EditProductView