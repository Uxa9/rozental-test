import { useLocation, useNavigate } from "react-router-dom"
import { Product } from "../../common/productStore/product"
import ProductForm from "../../components/productForm/ProductForm"
import { useEffect, useState } from "react";
import { connect } from "react-redux";
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore";
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher";


const EditProductView: React.FC<ProductStoreProps> = (props) => {
    
    const location = useLocation();
    const navigate = useNavigate();

    const [currentElement, setCurrentElement] = useState<Product>({id: 0, name: "", price: 0})

    useEffect(() => {
        let curEl = props.elements.find(el => el.id === parseInt(location.pathname.split('/').pop() || ""))!
        !curEl && navigate("/");
        setCurrentElement(curEl);
    }, [])

    const handleFormSubmit = (data: Product) => {
        props.editProduct({
            id: data.id,
            name: data.name,
            price: data.price,
        })

        navigate(`/product/${data.id}`)
    }

    return (
        <ProductForm 
            product={currentElement}
            resultHandler={handleFormSubmit} 
        />
    )
}

export default connect<StateProps, DispatcherProps>(
    productStoreState,
    productStoreDispatcher,
)(EditProductView)