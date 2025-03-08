import React, { useState } from 'react';
import Input from '../../atoms/Input/Input';
import Button from '../../atoms/Button/Button';
import './CitasForm.css';

interface CitaFormProps {
  onSubmit: (cita: CitaFormData) => void;
  initialData?: CitaFormData;
  onCancel: () => void;
}

interface CitaFormData {
  nombre: string;
  fecha: string;
  hora: string;
  motivo: string;
}

const CitaForm: React.FC<CitaFormProps> = ({ onSubmit, initialData, onCancel }) => {
  const [formData, setFormData] = useState<CitaFormData>(initialData || {
    nombre: '',
    fecha: '',
    hora: '',
    motivo: ''
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(formData);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  return (
    <form onSubmit={handleSubmit} className="cita-form">
      <div className="form-group">
        <Input
          label="Paciente"
          name="nombre"
          value={formData.nombre}
          onChange={handleChange}
          required
        />
      </div>
      <div className="form-group">
        <Input
          type="date"
          label="Fecha"
          name="fecha"
          value={formData.fecha}
          onChange={handleChange}
          required
        />
      </div>
      <div className="form-group">
        <Input
          type="time"
          label="Hora"
          name="hora"
          value={formData.hora}
          onChange={handleChange}
          required
        />
      </div>
      <div className="form-group">
        <Input
          label="Motivo"
          name="motivo"
          value={formData.motivo}
          onChange={handleChange}
          required
        />
      </div>
      <div className="form-actions">
        <Button
          text="Cancelar"
          variant="secondary"
          onClick={onCancel}
        />
        <Button
          text="Guardar"
          type="submit"
          variant="primary"
        />
      </div>
    </form>
  );
};

export default CitaForm;