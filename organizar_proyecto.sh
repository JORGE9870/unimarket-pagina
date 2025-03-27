#!/bin/bash

# Definir los cambios de nombres en carpetas y archivos
declare -A renombres=(
    ["conf"]="configuración"
    ["controllers"]="controladores"
    ["middleware/auth"]="middleware/autenticacion"
    ["middleware/error"]="middleware/errores"
    ["middleware/logging"]="middleware/registro"
    ["middleware/metrics"]="middleware/metricas"
    ["middleware/ratelimit"]="middleware/limitador"
    ["models"]="modelos"
    ["repositories"]="repositorios"
    ["routers"]="rutas"
    ["services"]="servicios"
    
    ["app.conf"]="aplicacion.conf"
    ["base_controller.go"]="base.go"
    ["branch.go"]="sucursal.go"
    ["business.go"]="negocio.go"
    ["cart.go"]="carrito.go"
    ["category.go"]="categoria.go"
    ["inventory_controller.go"]="controlador_inventario.go"
    ["notification.go"]="notificacion.go"
    ["order.go"]="pedido.go"
    ["payment.go"]="pago.go"
    ["permission.go"]="permiso.go"
    ["product_controller.go"]="controlador_productos.go"
    ["productos.go"]="productos.go"
    ["rating.go"]="calificacion.go"
    ["reseña de productos.go"]="reseña_productos.go"
    ["role.go"]="rol.go"
    ["store.go"]="tienda.go"
    ["user.go"]="usuario.go"
    ["product_repository.go"]="repositorio_productos.go"
    ["router.go"]="enrutador.go"
    
    ["inventory_service.go"]="servicio_inventario.go"
    ["product_service.go"]="servicio_productos.go"
    ["cart_service.go"]="servicio_carrito.go"
    ["category_service.go"]="servicio_categoria.go"
    ["notification_service.go"]="servicio_notificacion.go"
    ["order_service.go"]="servicio_pedidos.go"
    ["rating_service.go"]="servicio_calificacion.go"
    ["tienda_servicio.go"]="servicio_tienda.go"
    ["user_service.go"]="servicio_usuario.go"
)

# Renombrar carpetas
for key in "${!renombres[@]}"; do
    if [ -e "$key" ]; then
        mv "$key" "${renombres[$key]}"
    fi
done

echo "✅ Proyecto organizado con éxito 🚀🔥"
