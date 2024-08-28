import styles from "./CartView.module.scss";
import {logo} from "../../icons/index"
import { Black } from "../../components/images";
import { useNavigate } from "react-router-dom";
import Button from "../../components/button/Button";
import { Product } from "../../common/productStore/product";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, CartViewProps, StateProps } from "../../common/cart";
import CartDispatcher, { DispatcherProps } from "../../common/cart/dispatcher";
import { connect, useDispatch } from "react-redux";

type Props = {
    product: Product,
    amount: number
}

const ElementView: React.FC<Props> = (props) => {    

    const dispatch = useDispatch()

    const handleButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation()
        dispatch(CartDispatcher.removeFromCart(props.product.id))
    }

    return (
        <div className={styles["element-view"]}>
            <div>
                <img src={GetRandomSusById(props.product.id)} alt="product-image" className={styles["element-image"]}/>
                <div>
                    <p>{props.product.name}</p>
                    <p>{props.product.price} ₽</p>
                </div>
                <span>x{props.amount}</span>
            </div>
            <Button text="X" onClick={handleButtonClick}/>

        </div>
    )
}

export default ElementView