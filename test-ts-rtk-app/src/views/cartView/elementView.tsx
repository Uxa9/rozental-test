import styles from "./CartView.module.scss";
import {logo} from "../../icons/index"
import { Black } from "../../components/images";
import { useNavigate } from "react-router-dom";
import Button from "../../components/button/Button";
import { Product } from "../../common/productStore/product";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, CartViewProps, StateProps } from "../../common/cart";
import CartDispatcher, { DispatcherProps } from "../../common/cart/dispatcher";
import { connect } from "react-redux";


const ElementView: React.FC<CartViewProps> = (props) => {    
    const handleButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation()
        props.removeFromCart(props.own.product.id)
    }

    return (
        <div className={styles["element-view"]}>
            <div>
                <img src={GetRandomSusById(props.own.product.id)} alt="product-image" className={styles["element-image"]}/>
                <div>
                    <p>{props.own.product.name}</p>
                    <p>{props.own.product.price} ₽</p>
                </div>
                <span>x{props.own.amount}</span>
            </div>
            <Button text="X" onClick={handleButtonClick}/>

        </div>
    )
}

export default connect<StateProps, DispatcherProps>(
    cartStoreStateAndProps,
    CartDispatcher,
)(ElementView)