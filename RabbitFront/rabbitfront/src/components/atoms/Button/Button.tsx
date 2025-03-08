import React from 'react';
import './Button.css';

interface ButtonProps {
  text: string;
  type?: 'button' | 'submit' | 'reset';
  variant?: 'primary' | 'secondary' | 'danger';
  onClick?: () => void;
  className?: string;
}

const Button: React.FC<ButtonProps> = ({ 
  text, 
  type = 'button', 
  variant = 'primary', 
  onClick,
  className = '' 
}) => {
  return (
    <button
      type={type}
      className={`button button-${variant} ${className}`}
      onClick={onClick}
    >
      {text}
    </button>
  );
};

export default Button;