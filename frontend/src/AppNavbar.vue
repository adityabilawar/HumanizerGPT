<template>
    <nav class="navbar">
        <div class="navbar-brand">
            <router-link to="/" class="navbar-item">Home</router-link>
        </div>
        <div class="navbar-menu">
            <div class="navbar-end">
                <router-link to="/contact" class="navbar-item">Contact</router-link>
                <router-link to="/about" class="navbar-item">About</router-link>
            </div>
        </div>
    </nav>
</template>

<style scoped>
.navbar {
    background-color: #f5f5f5;
    padding: 1rem;
    font-family: Arial, sans-serif;  
}

.navbar-brand {
    font-weight: bold;
}

.navbar-item {
    margin-right: 1rem;
    color: #333;
    text-decoration: none;
}

.navbar-item:hover {
    color: #555;
}
</style>