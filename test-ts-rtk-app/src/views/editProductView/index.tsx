import { useLocation, useNavigate } from "react-router-dom"
import { Product } from "../../common/_productStore/product"
import ProductForm from "../../components/productForm/ProductForm"
import { useEffect, useState } from "react";
import { connect, useDispatch, useSelector } from "react-redux";
import { ProductStoreProps, StateProps, productStoreState } from "../../common/_productStore";
import productStoreDispatcher, { DispatcherProps } from "../../common/_productStore/dispatcher";
import { useAppDispatch } from "../..";
import { editProduct } from "../../common/productStore/slice";


const EditProductView = () => {
    
    const location = useLocation();
    const navigate = useNavigate();

    const {products} = useSelector((state: any) => state)
    const dispatch = useAppDispatch()

    const [currentElement, setCurrentElement] = useState<Product>({id: 0, name: "", price: 0})

    useEffect(() => {
        let curEl = products.find((el: Product) => el.id === parseInt(location.pathname.split('/').pop() || ""))!
        !curEl && navigate("/");
        setCurrentElement(curEl);
    }, [])

    const handleFormSubmit = (data: Product) => {

        dispatch(editProduct(data))

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