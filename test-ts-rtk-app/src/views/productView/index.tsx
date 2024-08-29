import { connect, useDispatch, useSelector } from "react-redux"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/_productStore"
import productStoreDispatcher, { DispatcherProps } from "../../common/_productStore/dispatcher"
import { useLocation, useNavigate } from "react-router-dom"
import { useEffect, useState } from "react"
import { Product } from "../../common/_productStore/product"
import Button from "../../components/button/Button"
import styles from "./ProductView.module.scss"
import { GetRandomSusById } from "../../randomSusAssert"
import { RootState, StoreDispathcer } from "../../common/indexStore"
import { useAppDispatch } from "../.."
import { deleteProduct } from "../../common/productStore/slice"



const ProductView = () => {

    const location = useLocation();    
    const navigate = useNavigate();

    const {products} = useSelector((state: RootState) => state)
    const dispatch = useAppDispatch()

    const [currentElement, setCurrentElement] = useState<Product>({id: 0, name: "", price: 0})

    useEffect(() =>{ 
        let curEl = products.find((el: Product) => el.id === parseInt(location.pathname.split('/').pop() || ""))!
        !curEl && navigate("/");
        setCurrentElement(curEl);
    }, [])

    const handleDeleteProduct = () => {
        dispatch(deleteProduct(currentElement.id))
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

export default ProductView