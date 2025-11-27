import { createBrowserRouter, RouterProvider, type RouteObject } from 'react-router';
import { UserInterface } from '../user-interface';
import { CalendarPage } from '../user-interface/pages/calendar';

const routeObject: RouteObject[] = [
  {
    path: '/',
    Component: UserInterface,
    children: [
      {
        path: '',
        Component: CalendarPage
      }
    ]
  },
]

export function Router() {
  const router = createRouter();

  return <RouterProvider router={router}></RouterProvider>;
}

function createRouter() {
  return createBrowserRouter(routeObject);
}
