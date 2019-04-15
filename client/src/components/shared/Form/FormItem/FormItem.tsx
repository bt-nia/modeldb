import * as React from 'react';

import styles from './FormItem.module.css';

interface IProps {
  label: string;
  additionalContent?: React.ReactNode;
  children: React.ReactNode;
}

const FormItem = ({ label, additionalContent, children }: IProps) => {
  return (
    <div className={styles.form_item}>
      <span className={styles.label}>{label}</span>
      {additionalContent && (
        <div className={styles.additional_content}>{additionalContent}</div>
      )}
      <div className={styles.content}>{children}</div>
    </div>
  );
};

export default FormItem;