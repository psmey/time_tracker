import Card from '@mui/joy/Card';
import { Outlet } from 'react-router';

export function Pages() {
  return (
    <Card style={{ flex: 1 }}>
      <Outlet />
    </Card>
  );
}
