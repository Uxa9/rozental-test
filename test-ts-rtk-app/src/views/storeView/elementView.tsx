import styles from "./StoreView.module.scss";
import { useNavigate } from "react-router-dom";
import Button from "../../components/button/Button";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, StateProps } from "../../common/_cart";
import CartDispatcher, { DispatcherProps } from "../../common/_cart/dispatcher";
import { connect, useDispatch, useSelector } from "react-redux";
import { Product } from "../../common/_productStore/product";
import { useAppDispatch } from "../..";
import { addToCart } from "../../common/cart/slice";

type Props = {
    product: Product
}

const ElementView: React.FC<Props> = (props) => {    
    const navigate = useNavigate()
    const dispatch = useAppDispatch()

    const { product } = props;
    
    const handleElementClick = () => {
        navigate(`/product/${product.id}`)
    }

    const handleButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation()
        dispatch(addToCart(product.id))
    }

    return (
        <div className={styles["element-view"]} onClick={handleElementClick}>
            <img src={GetRandomSusById(product.id)} alt="product-image" className={styles["element-image"]}/>
            <div>
                <p>{product.name}</p>
                <p>{product.price} ₽</p>
            </div>
            <Button text="Купить" onClick={handleButtonClick}/>

        </div>
    )
}

export default ElementView