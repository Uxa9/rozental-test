import { connect } from "react-redux"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/productStore"
import productStoreDispatcher, { DispatcherProps } from "../../common/productStore/dispatcher"
import { useLocation, useNavigate } from "react-router-dom"
import { useEffect, useState } from "react"
import { Product } from "../../common/productStore/product"
import Button from "../../components/button/Button"
import styles from "./ProductView.module.scss"
import { GetRandomSusById } from "../../randomSusAssert"



const ProductView: React.FC<ProductStoreProps> = (props) => {

    const location = useLocation();    
    const navigate = useNavigate();

    const [currentElement, setCurrentElement] = useState<Product>({id: 0, name: "", price: 0})

    useEffect(() =>{ 
        let curEl = props.elements.find(el => el.id === parseInt(location.pathname.split('/').pop() || ""))!
        !curEl && navigate("/");
        setCurrentElement(curEl);
    }, [])

    const handleDeleteProduct = () => {
        props.removeProduct(currentElement.id)
        navigate("/")
    }

    const handleEditProduct = () => {
        navigate(`/product/edit/${currentElement.id}`)
    }
    

    return (
        <div className={styles["product-card"]}>
            <img src={GetRandomSusById(currentElement.id)} alt="product-image"  />
            <p>
                {currentElement.name}
            </p>
            <p>
                {currentElement.price} ₽
            </p>
            <Button text="Удалить товар" onClick={handleDeleteProduct}/>
            <Button text="Редактировать" onClick={handleEditProduct}/>
        </div>
    )
}

export default connect<StateProps, DispatcherProps>(
    productStoreState,
    productStoreDispatcher,
)(ProductView)