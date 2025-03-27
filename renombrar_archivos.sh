#!/bin/bash

declare -A nombres=(
    ["base_controller.go"]="controlador_base.go"
    ["branch.go"]="sucursal.go"
    ["business.go"]="negocio.go"
    ["cart.go"]="carrito.go"
    ["category.go"]="categoria.go"
    ["delivery.go"]="entrega.go"
    ["inventory_controller.go"]="controlador_inventario.go"
    ["notification.go"]="notificacion.go"
    ["order.go"]="pedido.go"
    ["payment.go"]="pago.go"
    ["permission.go"]="permiso.go"
    ["product_controller.go"]="controlador_producto.go"
    ["rating.go"]="calificacion.go"
    ["role.go"]="rol.go"
    ["store.go"]="tienda.go"
    ["user.go"]="usuario.go"
    ["inventory_service.go"]="servicio_inventario.go"
    ["product_service.go"]="servicio_producto.go"
    ["cart_service.go"]="servicio_carrito.go"
    ["category_service.go"]="servicio_categoria.go"
    ["notification_service.go"]="servicio_notificacion.go"
    ["order_service.go"]="servicio_pedido.go"
    ["rating_service.go"]="servicio_calificacion.go"
    ["product_repository.go"]="repositorio_producto.go"
    ["router.go"]="enrutador.go"
)

# Renombrar archivos en todas las carpetas del proyecto
for archivo in "${!nombres[@]}"; do
    if find . -type f -name "$archivo" | grep -q .; then
        find . -type f -name "$archivo" -execdir mv {} "${nombres[$archivo]}" \;
        echo "Renombrado: $archivo -> ${nombres[$archivo]}"
    fi
done

# Renombrar archivos en la carpeta controladores
cd controladores

# Renombrar archivos que no siguen el patrón controlador_*.go
mv calificacion.go controlador_calificacion.go
mv carrito.go controlador_carrito.go
mv entrega.go controlador_entrega.go
mv notificacion.go controlador_notificacion.go
mv pago.go controlador_pago.go
mv permiso.go controlador_permiso.go
mv productos.go controlador_productos.go
mv rol.go controlador_rol.go
mv tienda.go controlador_tienda.go
mv usuario.go controlador_usuario.go

echo "✅ Todos los archivos han sido renombrados correctamente."
