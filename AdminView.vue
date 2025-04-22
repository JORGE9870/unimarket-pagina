<template>
  <div class="admin-view">
    <div class="admin-header">
      <h1>Panel de Administración</h1>
      <div class="admin-user-info" v-if="currentUser">
        <span>{{ currentUser.name }}</span>
        <button @click="logout" class="logout-btn">Cerrar sesión</button>
      </div>
    </div>

    <div v-if="!isAuthenticated" class="login-container">
      <div class="login-form">
        <h2>Acceso Administrador</h2>
        <div class="form-group">
          <label for="username">Usuario</label>
          <input 
            type="text" 
            id="username" 
            v-model="loginForm.username" 
            placeholder="Nombre de usuario"
          >
        </div>
        <div class="form-group">
          <label for="password">Contraseña</label>
          <input 
            type="password" 
            id="password" 
            v-model="loginForm.password" 
            placeholder="Contraseña"
          >
        </div>
        <div v-if="loginError" class="error-message">
          {{ loginError }}
        </div>
        <button @click="login" class="login-btn" :disabled="isLoggingIn">
          {{ isLoggingIn ? 'Iniciando sesión...' : 'Iniciar sesión' }}
        </button>
      </div>
    </div>

    <div v-else class="admin-dashboard">
      <div class="admin-sidebar">
        <nav>
          <button 
            @click="activeSection = 'dashboard'" 
            :class="{ active: activeSection === 'dashboard' }"
          >
            Dashboard
          </button>
          <button 
            @click="activeSection = 'anuncios'" 
            :class="{ active: activeSection === 'anuncios' }"
          >
            Anuncios
          </button>
          <button 
            @click="activeSection = 'categorias'" 
            :class="{ active: activeSection === 'categorias' }"
          >
            Categorías
          </button>
          <button 
            @click="activeSection = 'usuarios'" 
            :class="{ active: activeSection === 'usuarios' }"
          >
            Usuarios
          </button>
          <button 
            @click="activeSection = 'configuracion'" 
            :class="{ active: activeSection === 'configuracion' }"
          >
            Configuración
          </button>
        </nav>
      </div>

      <div class="admin-content">
        <Dashboard v-if="activeSection === 'dashboard'" />
        <Anuncios v-if="activeSection === 'anuncios'" />
        <Categorias v-if="activeSection === 'categorias'" />
        <Usuarios v-if="activeSection === 'usuarios'" />
        <Configuracion v-if="activeSection === 'configuracion'" />

        <!-- Lista de Usuarios -->
        <div class="admin-section">
          <h2>Usuarios</h2>
          <div class="search-bar">
            <input 
              type="text" 
              v-model="usuariosFilter" 
              placeholder="Buscar usuarios..."
            >
          </div>
          <div class="table-container">
            <table>
              <thead>
                <tr>
                  <th>Nombre</th>
                  <th>Email</th>
                  <th>Rol</th>
                  <th>Verificado</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="usuario in filteredUsuarios" :key="usuario.id">
                  <td>{{ usuario.nombreArtistico || usuario.name }}</td>
                  <td>{{ usuario.email }}</td>
                  <td>{{ usuario.role }}</td>
                  <td>
                    <span :class="['status-badge', usuario.verified ? 'verified' : 'pending']">
                      {{ usuario.verified ? 'Verificado' : 'Pendiente' }}
                    </span>
                  </td>
                  <td>
                    <button 
                      v-if="!usuario.verified" 
                      @click="verifyUsuario(usuario.id)" 
                      class="verify-btn"
                    >
                      Verificar
                    </button>
                    <button @click="editUsuario(usuario.id)" class="edit-btn">Editar</button>
                    <button @click="confirmDeleteUsuario(usuario.id)" class="delete-btn">Eliminar</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Dashboard from '@/components/admin/Dashboard.vue'
import Anuncios from '@/components/admin/Anuncios.vue'
import Categorias from '@/components/admin/Categorias.vue'
import Usuarios from '@/components/admin/Usuarios.vue'
import Configuracion from '@/components/admin/Configuracion.vue'

import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Router
const router = useRouter()

// Estado de autenticación
const isAuthenticated = ref(false)
const isLoggingIn = ref(false)
const loginError = ref('')
const loginForm = ref({
  username: '',
  password: ''
})

// Usuario actual
const currentUser = ref(null)

// Sección activa del panel
const activeSection = ref('dashboard')

// Datos de ejemplo para el dashboard
const stats = ref({
  totalAnuncios: 156,
  anunciosPendientes: 23,
  totalUsuarios: 87,
  visitasHoy: 1245
})

// Actividad reciente
const recentActivity = ref([
  { timestamp: new Date(Date.now() - 1000 * 60 * 5), description: 'Nuevo anuncio publicado: "Laura - Escort Premium"' },
  { timestamp: new Date(Date.now() - 1000 * 60 * 30), description: 'Usuario "maria89" ha actualizado su perfil' },
  { timestamp: new Date(Date.now() - 1000 * 60 * 60), description: 'Anuncio "Sofia VIP" ha sido aprobado por admin' },
  { timestamp: new Date(Date.now() - 1000 * 60 * 60 * 2), description: 'Nuevo usuario registrado: "carlos23"' },
  { timestamp: new Date(Date.now() - 1000 * 60 * 60 * 5), description: 'Categoría "Masajes" ha sido actualizada' }
])

// Datos de ejemplo para anuncios
const anuncios = ref([
  { 
    id: 1, 
    titulo: 'Sofía - Escort VIP', 
    descripcion: 'Servicios exclusivos de compañía', 
    ciudad: 'Madrid', 
    precio: '150€', 
    categoriaId: 1,
    estado: 'active',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 24 * 5)
  },
  { 
    id: 2, 
    titulo: 'Laura - Escort Premium', 
    descripcion: 'Experiencia inolvidable garantizada', 
    ciudad: 'Barcelona', 
    precio: '200€', 
    categoriaId: 1,
    estado: 'active',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 24 * 3)
  },
  { 
    id: 3, 
    titulo: 'Valentina - Escort de Lujo', 
    descripcion: 'Acompañante de alto nivel', 
    ciudad: 'Valencia', 
    precio: '180€', 
    categoriaId: 1,
    estado: 'pending',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 12)
  },
  { 
    id: 4, 
    titulo: 'Natalia - Escort Exclusiva', 
    descripcion: 'Atención personalizada', 
    ciudad: 'Sevilla', 
    precio: '160€', 
    categoriaId: 1,
    estado: 'active',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 24 * 7)
  },
  { 
    id: 5, 
    titulo: 'Centro de Masajes Oriental', 
    descripcion: 'Masajes relajantes y terapéuticos', 
    ciudad: 'Madrid', 
    precio: '80€', 
    categoriaId: 2,
    estado: 'active',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 24 * 10)
  },
  { 
    id: 6, 
    titulo: 'Masajes Tántricos', 
    descripcion: 'Experiencia única de relajación', 
    ciudad: 'Barcelona', 
    precio: '120€', 
    categoriaId: 2,
    estado: 'pending',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 6)
  },
  { 
    id: 7, 
    titulo: 'Carlos - Acompañante para Eventos', 
    descripcion: 'Elegante y culto para cualquier ocasión', 
    ciudad: 'Madrid', 
    precio: '200€', 
    categoriaId: 3,
    estado: 'rejected',
    fechaCreacion: new Date(Date.now() - 1000 * 60 * 60 * 24 * 2)
  }
])

// Filtros para anuncios
const anunciosFilter = ref('')
const anunciosFilterStatus = ref('all')
const anunciosPage = ref(1)
const anunciosPerPage = 5

// Datos de ejemplo para categorías
const categorias = ref([
  { id: 1, nombre: 'Escorts VIP', descripcion: 'Servicios exclusivos de compañía', iconoClass: 'fas fa-star' },
  { id: 2, nombre: 'Masajes', descripcion: 'Servicios de masajes profesionales', iconoClass: 'fas fa-spa' },
  { id: 3, nombre: 'Acompañantes', descripcion: 'Acompañantes para eventos y ocasiones especiales', iconoClass: 'fas fa-user-tie' },
  { id: 4, nombre: 'Servicios a domicilio', descripcion: 'Servicios que se desplazan a tu ubicación', iconoClass: 'fas fa-home' }
])

// Datos de ejemplo para usuarios
const usuarios = ref([
  { 
    id: 1, 
    nombre: 'Admin Principal', 
    email: 'admin@vipscort.com', 
    rol: 'admin', 
    activo: true,
    fechaRegistro: new Date(Date.now() - 1000 * 60 * 60 * 24 * 30)
  },
  { 
    id: 2, 
    nombre: 'Moderador 1', 
    email: 'mod1@vipscort.com', 
    rol: 'moderator', 
    activo: true,
    fechaRegistro: new Date(Date.now() - 1000 * 60 * 60 * 24 * 25)
  },
  { 
    id: 3, 
    nombre: 'María López', 
    email: 'maria@example.com', 
    rol: 'user', 
    activo: true,
    fechaRegistro: new Date(Date.now() - 1000 * 60 * 60 * 24 * 15)
  },
  { 
    id: 4, 
    nombre: 'Carlos Rodríguez', 
    email: 'carlos@example.com', 
    rol: 'user', 
    activo: false,
    fechaRegistro: new Date(Date.now() - 1000 * 60 * 60 * 24 * 10)
  },
  { 
    id: 5, 
    nombre: 'Ana Martínez', 
    email: 'ana@example.com', 
    rol: 'user', 
    activo: true,
    fechaRegistro: new Date(Date.now() - 1000 * 60 * 60 * 24 * 5)
  }
])

// Filtros para usuarios
const usuariosFilter = ref('')
const usuariosPage = ref(1)
const usuariosPerPage = 5

// Configuración del sitio
const siteConfig = ref({
  siteName: 'VIP-Scort',
  siteDescription: 'La mejor plataforma para encontrar servicios exclusivos de compañía',
  contactEmail: 'contacto@vipscort.com',
  contactPhone: '+34 612 345 678',
  maxImagesPerAd: 10,
  adDuration: 30,
  requireApproval: true
})

// Estado de modales
const showAddCategoryModal = ref(false)
const showAddUserModal = ref(false)
const showDeleteConfirmModal = ref(false)
const deleteConfirmMessage = ref('')
const deleteType = ref('')
const deleteId = ref(null)

// Formularios
const categoryForm = ref({
  nombre: '',
  descripcion: '',
  iconoClass: ''
})

const userForm = ref({
  nombre: '',
  email: '',
  password: '',
  rol: 'user',
  activo: true
})

// IDs para edición
const editingCategoryId = ref(null)
const editingUserId = ref(null)

// Computed properties
const filteredAnuncios = computed(() => {
  let filtered = anuncios.value

  // Filtrar por texto
  if (anunciosFilter.value) {
    const searchTerm = anunciosFilter.value.toLowerCase()
    filtered = filtered.filter(a => 
      a.titulo.toLowerCase().includes(searchTerm) || 
      a.descripcion.toLowerCase().includes(searchTerm) || 
      a.ciudad.toLowerCase().includes(searchTerm)
    )
  }

  // Filtrar por estado
  if (anunciosFilterStatus.value !== 'all') {
    filtered = filtered.filter(a => a.estado === anunciosFilterStatus.value)
  }

  // Paginación
  const start = (anunciosPage.value - 1) * anunciosPerPage
  const end = start + anunciosPerPage
  return filtered.slice(start, end)
})

const totalAnunciosPages = computed(() => {
  let filtered = anuncios.value

  // Filtrar por texto
  if (anunciosFilter.value) {
    const searchTerm = anunciosFilter.value.toLowerCase()
    filtered = filtered.filter(a => 
      a.titulo.toLowerCase().includes(searchTerm) || 
      a.descripcion.toLowerCase().includes(searchTerm) || 
      a.ciudad.toLowerCase().includes(searchTerm)
    )
  }

  // Filtrar por estado
  if (anunciosFilterStatus.value !== 'all') {
    filtered = filtered.filter(a => a.estado === anunciosFilterStatus.value)
  }

  return Math.ceil(filtered.length / anunciosPerPage)
})

const filteredUsuarios = computed(() => {
  let filtered = usuarios.value

  // Filtrar por texto
  if (usuariosFilter.value) {
    const searchTerm = usuariosFilter.value.toLowerCase()
    filtered = filtered.filter(u => 
      u.nombre.toLowerCase().includes(searchTerm) || 
      u.email.toLowerCase().includes(searchTerm)
    )
  }

  // Paginación
  const start = (usuariosPage.value - 1) * usuariosPerPage
  const end = start + usuariosPerPage
  return filtered.slice(start, end)
})

const totalUsuariosPages = computed(() => {
  let filtered = usuarios.value

  // Filtrar por texto
  if (usuariosFilter.value) {
    const searchTerm = usuariosFilter.value.toLowerCase()
    filtered = filtered.filter(u => 
      u.nombre.toLowerCase().includes(searchTerm) || 
      u.email.toLowerCase().includes(searchTerm)
    )
  }

  return Math.ceil(filtered.length / usuariosPerPage)
})

// Funciones
function formatDate(date) {
  const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' }
  return new Date(date).toLocaleDateString('es-ES', options)
}

function getCategoryName(id) {
  const category = categorias.value.find(c => c.id === id)
  return category ? category.nombre : 'Desconocida'
}

function getStatusText(status) {
  switch (status) {
    case 'active':
      return 'Activo'
    case 'pending':
      return 'Pendiente'
    case 'rejected':
      return 'Rechazado'
    default:
      return 'Desconocido'
  }
}

function getAnunciosByCategoryCount(categoryId) {
  return anuncios.value.filter(a => a.categoriaId === categoryId).length
}

function login() {
  isLoggingIn.value = true
  loginError.value = ''

  // Simulación de autenticación
  setTimeout(() => {
    if (loginForm.value.username === 'admin' && loginForm.value.password === 'admin') {
      isAuthenticated.value = true
      currentUser.value = { id: 1, name: 'Admin Principal' }
    } else {
      loginError.value = 'Usuario o contraseña incorrectos'
    }
    isLoggingIn.value = false
  }, 1000)
}

function logout() {
  isAuthenticated.value = false
  currentUser.value = null
}

function viewAnuncio(id) {
  console.log('Ver anuncio', id)
}

function editAnuncio(id) {
  console.log('Editar anuncio', id)
}

function confirmDeleteAnuncio(id) {
  deleteConfirmMessage.value = '¿Estás seguro de que deseas eliminar este anuncio?'
  deleteType.value = 'anuncio'
  deleteId.value = id
  showDeleteConfirmModal.value = true
}

function approveAnuncio(id) {
  const anuncio = anuncios.value.find(a => a.id === id)
  if (anuncio) {
    anuncio.estado = 'active'
  }
}

function rejectAnuncio(id) {
  const anuncio = anuncios.value.find(a => a.id === id)
  if (anuncio) {
    anuncio.estado = 'rejected'
  }
}

function editCategoria(id) {
  const categoria = categorias.value.find(c => c.id === id)
  if (categoria) {
    categoryForm.value = { ...categoria }
    editingCategoryId.value = id
    showAddCategoryModal.value = true
  }
}

function saveCategory() {
  if (editingCategoryId.value) {
    const index = categorias.value.findIndex(c => c.id === editingCategoryId.value)
    if (index !== -1) {
      categorias.value[index] = { ...categoryForm.value, id: editingCategoryId.value }
    }
  } else {
    const newCategory = { ...categoryForm.value, id: Date.now() }
    categorias.value.push(newCategory)
  }
  closeAddCategoryModal()
}

function closeAddCategoryModal() {
  showAddCategoryModal.value = false
  categoryForm.value = { nombre: '', descripcion: '', iconoClass: '' }
  editingCategoryId.value = null
}

function confirmDeleteCategoria(id) {
  deleteConfirmMessage.value = '¿Estás seguro de que deseas eliminar esta categoría?'
  deleteType.value = 'categoria'
  deleteId.value = id
  showDeleteConfirmModal.value = true
}

function editUsuario(id) {
  const usuario = usuarios.value.find(u => u.id === id)
  if (usuario) {
    userForm.value = { ...usuario, password: '' }
    editingUserId.value = id
    showAddUserModal.value = true
  }
}

function saveUser() {
  if (editingUserId.value) {
    const index = usuarios.value.findIndex(u => u.id === editingUserId.value)
    if (index !== -1) {
      usuarios.value[index] = { ...userForm.value, id: editingUserId.value }
    }
  } else {
    const newUser = { ...userForm.value, id: Date.now() }
    usuarios.value.push(newUser)
  }
  closeAddUserModal()
}

function closeAddUserModal() {
  showAddUserModal.value = false
  userForm.value = { nombre: '', email: '', password: '', rol: 'user', activo: true }
  editingUserId.value = null
}

function toggleUsuarioStatus(id) {
  const usuario = usuarios.value.find(u => u.id === id)
  if (usuario) {
    usuario.activo = !usuario.activo
  }
}

function confirmDeleteUsuario(id) {
  deleteConfirmMessage.value = '¿Estás seguro de que deseas eliminar este usuario?'
  deleteType.value = 'usuario'
  deleteId.value = id
  showDeleteConfirmModal.value = true
}

function cancelDelete() {
  showDeleteConfirmModal.value = false
  deleteConfirmMessage.value = ''
  deleteType.value = ''
  deleteId.value = null
}

function confirmDelete() {
  if (deleteType.value === 'anuncio') {
    const index = anuncios.value.findIndex(a => a.id === deleteId.value)
    if (index !== -1) {
      anuncios.value.splice(index, 1)
    }
  } else if (deleteType.value === 'categoria') {
    const index = categorias.value.findIndex(c => c.id === deleteId.value)
    if (index !== -1) {
      categorias.value.splice(index, 1)
    }
  } else if (deleteType.value === 'usuario') {
    const index = usuarios.value.findIndex(u => u.id === deleteId.value)
    if (index !== -1) {
      usuarios.value.splice(index, 1)
    }
  }
  cancelDelete()
}

function saveConfiguration() {
  console.log('Configuración guardada', siteConfig.value)
}

const verifyUsuario = async (id) => {
  try {
    const authStore = useAuthStore()
    await authStore.verifyUser(id)
    
    // Actualizar la lista de usuarios
    const index = usuarios.value.findIndex(u => u.id === id)
    if (index !== -1) {
      usuarios.value[index].verified = true
    }
  } catch (error) {
    console.error('Error al verificar usuario:', error)
    alert('Error al verificar el usuario. Por favor, inténtalo de nuevo.')
  }
}
</script>
