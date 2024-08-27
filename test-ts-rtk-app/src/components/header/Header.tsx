import styles from './Header.module.scss'
import {logo, cart} from '../../icons/index'
import { useNavigate } from 'react-router-dom'


const Header = () => {

    const navigate = useNavigate();

    const handleClick = (path: string) => {
        navigate(path)
    }

    return (
        <header className={styles.header}>
            <div className={styles["company-info"]}>
                <img src={logo} className={styles["app-logo"]} alt="logo" onClick={() => handleClick("/")} />
                <span>Test store</span>
                <span className={styles["active-button"]} onClick={() => handleClick("/product/add")}>Add product</span>
            </div>
            <img src={cart} className={styles["cart-logo"]} alt="cart" onClick={() => handleClick("/cart")}/>
        </header>
    )
}

export default Header