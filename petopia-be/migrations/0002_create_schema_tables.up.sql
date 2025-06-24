-- Create User Accounts Table
CREATE TABLE user_accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create User Addresses Table
CREATE TABLE user_addresses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user_accounts(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    pincode VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create User Card Details Table
CREATE TABLE user_card_details (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user_accounts(id) ON DELETE CASCADE,
    card_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Seller Info Table
CREATE TABLE seller_info (
    id SERIAL PRIMARY KEY,
    seller_name VARCHAR(255) NOT NULL,
    seller_info TEXT,
    seller_address TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Products Table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Inventory Information Table
CREATE TABLE inventory_information (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    seller_id INTEGER NOT NULL REFERENCES seller_info(id) ON DELETE CASCADE,
    inventory_location TEXT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Cart Table
CREATE TABLE cart (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user_accounts(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Orders Table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user_accounts(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    total_price NUMERIC(10, 2) NOT NULL CHECK (total_price >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Products Ordered Table
CREATE TABLE products_ordered (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    delivery_status VARCHAR(20) NOT NULL CHECK (delivery_status IN ('Not Delivered', 'Delivered')) DEFAULT 'Not Delivered',
    return_status VARCHAR(20) NOT NULL CHECK (return_status IN ('Return Requested', 'Returned', 'None')) DEFAULT 'None',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create Shipping Table
CREATE TABLE shipping (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user_accounts(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    location_tracking JSONB,
    assigned_delivery_partner VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL REFERENCES user_accounts(id),
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by INTEGER REFERENCES user_accounts(id),
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted_by INTEGER REFERENCES user_accounts(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);