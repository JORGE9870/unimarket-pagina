<template>
  <div class="admin-dashboard">
    <header class="admin-header">
      <div class="container">
        <div class="header-content">
          <h1 class="dashboard-title">Panel de Administración</h1>
          <div class="admin-actions">
            <button class="action-button" @click="showUserManagement">
              <span class="icon">👥</span>
              Gestión de Usuarios
            </button>
            <button class="action-button" @click="showBusinessManagement">
              <span class="icon">🏪</span>
              Gestión de Negocios
            </button>
            <button class="action-button" @click="showReports">
              <span class="icon">📊</span>
              Reportes
            </button>
          </div>
        </div>
      </div>
    </header>

    <main class="admin-content">
      <div class="container">
        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon">👥</div>
            <div class="stat-info">
              <h3>Usuarios Totales</h3>
              <p class="stat-number">{{ totalUsers }}</p>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">🏪</div>
            <div class="stat-info">
              <h3>Negocios Registrados</h3>
              <p class="stat-number">{{ totalBusinesses }}</p>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">💰</div>
            <div class="stat-info">
              <h3>Ingresos Totales</h3>
              <p class="stat-number">${{ totalRevenue }}</p>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">📦</div>
            <div class="stat-info">
              <h3>Pedidos del Día</h3>
              <p class="stat-number">{{ dailyOrders }}</p>
            </div>
          </div>
        </div>

        <div class="management-section">
          <h2 class="section-title">Actividad Reciente</h2>
          <div class="activity-list">
            <div v-for="activity in recentActivities" :key="activity.id" class="activity-item">
              <div class="activity-icon">{{ activity.icon }}</div>
              <div class="activity-details">
                <p class="activity-description">{{ activity.description }}</p>
                <span class="activity-time">{{ activity.time }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const totalUsers = ref(1250)
const totalBusinesses = ref(350)
const totalRevenue = ref('25,000')
const dailyOrders = ref(180)

const recentActivities = ref([
  {
    id: 1,
    icon: '👤',
    description: 'Nuevo usuario registrado: Juan Pérez',
    time: 'Hace 5 minutos'
  },
  {
    id: 2,
    icon: '🏪',
    description: 'Nuevo negocio aprobado: Restaurante La Casa',
    time: 'Hace 15 minutos'
  },
  {
    id: 3,
    icon: '💰',
    description: 'Pago procesado: $1,200',
    time: 'Hace 30 minutos'
  },
  {
    id: 4,
    icon: '📦',
    description: 'Pedido completado: #12345',
    time: 'Hace 1 hora'
  }
])

const showUserManagement = () => {
  // Implementar lógica de gestión de usuarios
}

const showBusinessManagement = () => {
  // Implementar lógica de gestión de negocios
}

const showReports = () => {
  // Implementar lógica de reportes
}
</script>

<style lang="scss" scoped>
.admin-dashboard {
  min-height: 100vh;
  background: #f8f9fa;

  .admin-header {
    background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
    padding: 2rem 0;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);

    .header-content {
      display: flex;
      justify-content: space-between;
      align-items: center;

      .dashboard-title {
        color: #ffffff;
        font-size: 2rem;
        font-weight: 600;
        letter-spacing: 0.5px;
      }

      .admin-actions {
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

  .admin-content {
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

    .management-section {
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

      .activity-list {
        display: flex;
        flex-direction: column;
        gap: 1rem;

        .activity-item {
          display: flex;
          align-items: center;
          gap: 1rem;
          padding: 1rem;
          border-radius: 8px;
          background: #f8f9fa;
          transition: all 0.3s ease;

          &:hover {
            background: #e9ecef;
          }

          .activity-icon {
            font-size: 1.5rem;
            background: #ffffff;
            width: 50px;
            height: 50px;
            display: flex;
            align-items: center;
            justify-content: center;
            border-radius: 12px;
            box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
          }

          .activity-details {
            flex: 1;

            .activity-description {
              font-size: 1rem;
              color: #1a1a2e;
              margin-bottom: 0.25rem;
            }

            .activity-time {
              font-size: 0.875rem;
              color: #6c757d;
            }
          }
        }
      }
    }
  }
}
</style> 