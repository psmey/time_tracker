import { Pages } from "./pages";
import { SideBar } from "./side-bar";

export function UserInterface() {
  return (
    <>
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          gap: '1rem',
          height: '100%',
          padding: '2rem',
          boxSizing: 'border-box',
        }}
      >
        <h3>⏱️ <b><i>Tempo</i> Time Tracker</b></h3>
        <div style={{
          display: 'flex',
          flexDirection: 'row',
          height: '100%',
          gap: '1rem',
        }}>
          <SideBar></SideBar>
          <Pages></Pages>
        </div>
      </div >
    </>
  );
}
