<template>
  <div class="business-dashboard">
    <header class="business-header">
      <div class="container">
        <div class="header-content">
          <div class="business-info">
            <h1 class="business-name">{{ businessName }}</h1>
            <p class="business-status">Estado: <span :class="statusClass">{{ businessStatus }}</span></p>
          </div>
          <div class="business-actions">
            <button class="action-button" @click="showOrders">
              <span class="icon">📦</span>
              Pedidos
            </button>
            <button class="action-button" @click="showMenu">
              <span class="icon">🍽️</span>
              Menú
            </button>
            <button class="action-button" @click="showAnalytics">
              <span class="icon">📊</span>
              Análisis
            </button>
          </div>
        </div>
      </div>
    </header>

    <main class="business-content">
      <div class="container">
        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon">📦</div>
            <div class="stat-info">
              <h3>Pedidos Hoy</h3>
              <p class="stat-number">{{ todayOrders }}</p>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">💰</div>
            <div class="stat-info">
              <h3>Ingresos Hoy</h3>
              <p class="stat-number">${{ todayRevenue }}</p>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">⭐</div>
            <div class="stat-info">
              <h3>Calificación</h3>
              <p class="stat-number">{{ rating }}/5</p>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">👥</div>
            <div class="stat-info">
              <h3>Empleados</h3>
              <p class="stat-number">{{ totalEmployees }}</p>
            </div>
          </div>
        </div>

        <div class="orders-section">
          <h2 class="section-title">Pedidos Recientes</h2>
          <div class="orders-list">
            <div v-for="order in recentOrders" :key="order.id" class="order-card">
              <div class="order-header">
                <span class="order-id">#{{ order.id }}</span>
                <span class="order-status" :class="order.status">{{ order.status }}</span>
              </div>
              <div class="order-details">
                <p class="order-items">{{ order.items }}</p>
                <p class="order-total">Total: ${{ order.total }}</p>
              </div>
              <div class="order-footer">
                <span class="order-time">{{ order.time }}</span>
                <button class="action-button" @click="viewOrder(order.id)">
                  Ver Detalles
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const businessName = ref('Mi Restaurante')
const businessStatus = ref('Activo')
const statusClass = computed(() => ({
  'status-active': businessStatus.value === 'Activo',
  'status-inactive': businessStatus.value === 'Inactivo'
}))

const todayOrders = ref(25)
const todayRevenue = ref('1,250')
const rating = ref(4.5)
const totalEmployees = ref(8)

const recentOrders = ref([
  {
    id: '12345',
    status: 'En preparación',
    items: '2 Hamburguesas, 1 Papas fritas',
    total: '45.99',
    time: 'Hace 15 minutos'
  },
  {
    id: '12344',
    status: 'Completado',
    items: '1 Pizza familiar, 2 Refrescos',
    total: '35.50',
    time: 'Hace 30 minutos'
  },
  {
    id: '12343',
    status: 'Pendiente',
    items: '3 Tacos, 1 Ensalada',
    total: '28.75',
    time: 'Hace 45 minutos'
  }
])

const showOrders = () => {
  // Implementar lógica de visualización de pedidos
}

const showMenu = () => {
  // Implementar lógica de gestión del menú
}

const showAnalytics = () => {
  // Implementar lógica de análisis
}

const viewOrder = (orderId) => {
  // Implementar lógica de visualización de detalles del pedido
}
</script>

<style lang="scss" scoped>
.business-dashboard {
  min-height: 100vh;
  background: #f8f9fa;

  .business-header {
    background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
    padding: 2rem 0;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);

    .header-content {
      display: flex;
      justify-content: space-between;
      align-items: center;

      .business-info {
        .business-name {
          color: #ffffff;
          font-size: 2rem;
          font-weight: 600;
          letter-spacing: 0.5px;
          margin-bottom: 0.5rem;
        }

        .business-status {
          color: rgba(255, 255, 255, 0.8);
          font-size: 1rem;

          .status-active {
            color: #4caf50;
          }

          .status-inactive {
            color: #f44336;
          }
        }
      }

      .business-actions {
        display: flex;
        gap: 1rem;

        .action-button {
          display: flex;
          align-items: center;
          gap: 0.5rem;
          padding: 0.75rem 1.5rem;
          background: rgba(255, 255, 255, 0.1);
          border: 1px solid rgba(255, 255, 255, 0.2);
          border-radius: 8px;
          color: #ffffff;
          font-size: 1rem;
          cursor: pointer;
          transition: all 0.3s ease;

          &:hover {
            background: rgba(255, 255, 255, 0.2);
            transform: translateY(-2px);
          }

          .icon {
            font-size: 1.25rem;
          }
        }
      }
    }
  }

  .business-content {
    padding: 3rem 0;

    .stats-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 2rem;
      margin-bottom: 3rem;

      .stat-card {
        background: #ffffff;
        padding: 2rem;
        border-radius: 12px;
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
        display: flex;
        align-items: center;
        gap: 1.5rem;
        transition: all 0.3s ease;

        &:hover {
          transform: translateY(-5px);
          box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
        }

        .stat-icon {
          font-size: 2.5rem;
          background: #f8f9fa;
          width: 80px;
          height: 80px;
          display: flex;
          align-items: center;
          justify-content: center;
          border-radius: 16px;
        }

        .stat-info {
          h3 {
            font-size: 1rem;
            color: #6c757d;
            margin-bottom: 0.5rem;
          }

          .stat-number {
            font-size: 1.75rem;
            color: #1a1a2e;
            font-weight: 600;
          }
        }
      }
    }

    .orders-section {
      background: #ffffff;
      padding: 2rem;
      border-radius: 12px;
      box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);

      .section-title {
        font-size: 1.5rem;
        color: #1a1a2e;
        margin-bottom: 1.5rem;
        font-weight: 600;
      }

      .orders-list {
        display: flex;
        flex-direction: column;
        gap: 1rem;

        .order-card {
          background: #f8f9fa;
          padding: 1.5rem;
          border-radius: 12px;
          transition: all 0.3s ease;

          &:hover {
            background: #e9ecef;
          }

          .order-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 1rem;

            .order-id {
              font-size: 1.1rem;
              color: #1a1a2e;
              font-weight: 600;
            }

            .order-status {
              padding: 0.25rem 0.75rem;
              border-radius: 20px;
              font-size: 0.875rem;
              font-weight: 500;

              &.En-preparación {
                background: #fff3e0;
                color: #e65100;
              }

              &.Completado {
                background: #e8f5e9;
                color: #2e7d32;
              }

              &.Pendiente {
                background: #e3f2fd;
                color: #1565c0;
              }
            }
          }

          .order-details {
            margin-bottom: 1rem;

            .order-items {
              font-size: 1rem;
              color: #1a1a2e;
              margin-bottom: 0.5rem;
            }

            .order-total {
              font-size: 1.1rem;
              color: #1a1a2e;
              font-weight: 600;
            }
          }

          .order-footer {
            display: flex;
            justify-content: space-between;
            align-items: center;

            .order-time {
              font-size: 0.875rem;
              color: #6c757d;
            }

            .action-button {
              padding: 0.5rem 1rem;
              background: #1a1a2e;
              color: #ffffff;
              border: none;
              border-radius: 6px;
              cursor: pointer;
              transition: all 0.3s ease;

              &:hover {
                background: #16213e;
                transform: translateY(-2px);
              }
            }
          }
        }
      }
    }
  }
}
</style> 