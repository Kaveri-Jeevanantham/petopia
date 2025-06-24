import { useEffect, useState } from 'react';
import ProductList from '../components/ProductList';

const ProductListPage = () => {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        // Fetch products from the API
        fetch('http://localhost:8080/api/products')
            .then(response => response.json())
            .then(data => setProducts(data))
            .catch(error => console.error('Error fetching products:', error));
    }, []);

    return (
        <div>
            <h1>Product List</h1>
            <ProductList products={products} />
        </div>
    );
};

export default ProductListPage;
