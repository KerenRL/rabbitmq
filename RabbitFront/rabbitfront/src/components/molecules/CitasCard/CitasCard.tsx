import React from 'react';
import './CitasCard.css';

interface CitaCardProps {
  id: string;
  nombre: string;
  fecha: string;
  hora: string;
  motivo: string;
  onEdit: (id: string) => void;
  onDelete: (id: string) => void;
  onView: (id: string) => void;
}

const CitaCard: React.FC<CitaCardProps> = ({ id, nombre, fecha, hora, motivo, onEdit, onDelete, onView }) => {
  return (
    <div className="cita-card">
      <h3>{nombre}</h3>
      <p>{fecha} a las {hora}</p>
      <p>{motivo}</p>
      <div className="cita-card-actions">
        <button onClick={() => onView(id)}>Ver</button>
        <button onClick={() => onEdit(id)}>Editar</button>
        <button onClick={() => onDelete(id)}>Eliminar</button>
      </div>
    </div>
  );
};

export default CitaCard;