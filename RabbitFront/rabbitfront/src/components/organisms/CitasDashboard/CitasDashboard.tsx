import React, { useState, useEffect } from 'react';
import CitaCard from '../../molecules/CitasCard/CitasCard';
import Button from '../../atoms/Button/Button';
import Modal from '../../molecules/Modal/Modal';
import CitaForm from '../../molecules/CitasForm/CitasForm';
import websocketService from '../../../services/websocketServices';
import './CitasDashboard.css';

interface Cita {
  id: string;
  nombre: string;
  fecha: string;
  hora: string;
  motivo: string;
}

const CitasDashboard: React.FC = () => {
  const [citas, setCitas] = useState<Cita[]>([]);
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [selectedCita, setSelectedCita] = useState<Cita | null>(null);
  const [modalType, setModalType] = useState<'create' | 'edit' | 'view'>('create');
  const [notifications, setNotifications] = useState<string[]>([]);

  useEffect(() => {
    fetchCitas();

    // Conectar al WebSocket
    websocketService.connect('ws://localhost:8082/ws');

    // Agregar listener para recibir notificaciones
    const handleNotification = (message: string) => {
      setNotifications((prev) => [...prev, message]);
    };
    websocketService.addListener(handleNotification);

    // Limpiar el listener al desmontar el componente
    return () => {
      websocketService.removeListener(handleNotification);
    };
  }, []);

  const fetchCitas = async () => {
    try {
      const response = await fetch('http://localhost:8080/citas');
      const data = await response.json();
      setCitas(data);
    } catch (error) {
      console.error('Error al obtener citas:', error);
    }
  };

  const handleCreateCita = () => {
    setModalType('create');
    setSelectedCita(null);
    setIsFormOpen(true);
  };

  const handleEditCita = (id: string) => {
    const cita = citas.find(c => c.id === id);
    if (cita) {
      setSelectedCita(cita);
      setModalType('edit');
      setIsFormOpen(true);
    }
  };

  const handleViewCita = (id: string) => {
    const cita = citas.find(c => c.id === id);
    if (cita) {
      setSelectedCita(cita);
      setModalType('view');
      setIsFormOpen(true);
    }
  };

  const handleDeleteCita = async (id: string) => {
    if (window.confirm('¿Estás seguro de que deseas eliminar esta cita?')) {
      try {
        await fetch(`http://localhost:8080/citas/${id}`, {
          method: 'DELETE'
        });
        setCitas(citas.filter(cita => cita.id !== id));
      } catch (error) {
        console.error('Error al eliminar cita:', error);
      }
    }
  };

  const handleSubmit = async (formData: Omit<Cita, 'id'>) => {
    try {
      if (modalType === 'edit' && selectedCita) {
        await fetch(`http://localhost:8080/citas/${selectedCita.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(formData)
        });
      } else {
        await fetch('http://localhost:8080/citas', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(formData)
        });
      }
      await fetchCitas();
      setIsFormOpen(false);
    } catch (error) {
      console.error('Error al guardar cita:', error);
    }
  };

  return (
    <div className="dashboard-container">
      <div className="dashboard-header">
        <h1 className="dashboard-title">Gestión de Citas</h1>
        <Button 
          text="Nueva Cita" 
          onClick={handleCreateCita}
          className="create-button"
        />
      </div>

      <div className="dashboard-grid">
        {citas.length > 0 ? (
          citas.map(cita => (
            <CitaCard
              key={cita.id}
              {...cita}
              onEdit={handleEditCita}
              onDelete={handleDeleteCita}
              onView={handleViewCita}
            />
          ))
        ) : (
          <div className="empty-state">
            <p className="empty-state-text">No hay citas registradas</p>
            <Button 
              text="Crear Primera Cita" 
              onClick={handleCreateCita}
            />
          </div>
        )}
      </div>

      <div className="notifications">
        <h2>Notificaciones</h2>
        <ul>
          {notifications.map((notification, index) => (
            <li key={index}>{notification}</li>
          ))}
        </ul>
      </div>

      <Modal
        isOpen={isFormOpen}
        onClose={() => setIsFormOpen(false)}
        title={modalType === 'create' ? 'Nueva Cita' : modalType === 'edit' ? 'Editar Cita' : 'Ver Cita'}
      >
        <CitaForm
          onSubmit={handleSubmit}
          initialData={selectedCita || undefined}
          onCancel={() => setIsFormOpen(false)}
        />
      </Modal>
    </div>
  );
};

export default CitasDashboard;