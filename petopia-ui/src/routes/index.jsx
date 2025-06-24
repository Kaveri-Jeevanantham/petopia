import { Route, Routes } from 'react-router-dom';
import ProductListPage from '../pages/ProductListPage';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/products" element={<ProductListPage />} />
            {/* Add more routes here as needed */}
        </Routes>
    );
};

export default AppRoutes;
