#!/bin/bash

declare -A duplicados=(
    ["controlador_producto.go"]="productos.go"
    ["controlador_productos.go"]="productos.go"
    ["controlador_tienda.go"]="tienda.go"
    ["reseña de productos.go"]="productos.go"
    ["inventario.go"]="controlador_inventario.go"
    ["negocio.go"]="sucursal.go"
    ["sucursal.go"]="rol.go"
    ["pedido.go"]="notificacion.go"
    ["categoria.go"]="permiso.go"
)

# Eliminar archivos duplicados
for archivo in "${!duplicados[@]}"; do
    if [ -f "controladores/$archivo" ]; then
        echo "🗑️ Eliminando duplicado: controladores/$archivo"
        rm "controladores/$archivo"
    fi
done

echo "✅ Archivos duplicados eliminados."
