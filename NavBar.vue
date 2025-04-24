<template>
  <nav class="navbar">
    <div class="navbar-container">
      <router-link to="/" class="logo">
        <img src="@/assets/logo.png" alt="Unimarka" class="logo-image" />
      </router-link>
      <div class="nav-links">
        <router-link to="/" class="nav-link">Inicio</router-link>
        <router-link to="/pedidos" class="nav-link">Pedidos</router-link>
        <router-link to="/categorias" class="nav-link">Categorías</router-link>
        <router-link to="/servicios" class="nav-link">Servicios y productos</router-link>
        <router-link to="/proposito" class="nav-link">Nuestro Propósito</router-link>
        <router-link to="/aliados" class="nav-link">Aliados Comerciales</router-link>
        <router-link to="/portales" class="nav-link">Portales</router-link>
        <router-link to="/contacto" class="nav-link">Contacto</router-link>
      </div>
      <div class="nav-actions">
        <button class="action-button">
          Actualizar
          <span class="icon">🔄</span>
        </button>
        <button class="menu-button">
          <span class="icon">⋮</span>
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup>
// No se requiere lógica adicional por ahora
</script>

<style scoped>
.navbar {
  background: linear-gradient(135deg, #2c3e50 0%, #3498db 100%);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.navbar-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  display: flex;
  align-items: center;
  text-decoration: none;
}

.logo-image {
  height: 40px;
  width: auto;
  filter: brightness(0) invert(1);
}

.nav-links {
  display: flex;
  gap: 2rem;
  align-items: center;
}

.nav-link {
  color: rgba(255, 255, 255, 0.9);
  text-decoration: none;
  font-size: 0.95rem;
  font-weight: 500;
  transition: all 0.3s ease;
  padding: 0.5rem 1rem;
  border-radius: 6px;
}

.nav-link:hover {
  color: #ffffff;
  background: rgba(255, 255, 255, 0.1);
}

.nav-link.router-link-active {
  color: #ffffff;
  background: rgba(255, 255, 255, 0.15);
  font-weight: 600;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.action-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1.25rem;
  background-color: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 6px;
  color: #ffffff;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.action-button:hover {
  background-color: rgba(255, 255, 255, 0.2);
  transform: translateY(-1px);
}

.menu-button {
  background: none;
  border: none;
  font-size: 1.5rem;
  color: rgba(255, 255, 255, 0.9);
  cursor: pointer;
  padding: 0.25rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.menu-button:hover {
  color: #ffffff;
  transform: rotate(90deg);
}

.icon {
  font-size: 1.1rem;
}

@media (max-width: 1024px) {
  .nav-links {
    display: none;
  }
  
  .navbar-container {
    padding: 0.75rem 1rem;
  }
  
  .action-button {
    padding: 0.5rem 1rem;
  }
}
</style> 