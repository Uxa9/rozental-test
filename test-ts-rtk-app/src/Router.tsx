import { Navigate, Route, Routes } from "react-router-dom"
import MainLayout from "./components/layouts/MainLayout"
import ProductView from "./views/productView"
import StoreView from "./views/storeView"
import CartView from "./views/cartView"
import AddProductView from "./views/addProductView"
import EditProductView from "./views/editProductView"

const AppRouter = () => {

    return (
        <MainLayout>
            <Routes>
                    <Route path="" element={<StoreView />} />
                    <Route path="product">
                        <Route path=":id" element={<ProductView />} />
                        <Route path="add" element={<AddProductView />} />
                        <Route path="edit/:id" element={<EditProductView />}/>
                    </Route>
                    <Route path="/cart" element={<CartView />} />
                    <Route path="*" element={<Navigate to="" replace={true} />}/>
            </Routes>
        </MainLayout>
    )
}

export default AppRouter