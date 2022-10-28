import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { RoomInterface } from "../interfaces/IRoom";
import { GetRoom } from "../services/HttpClientService";


function Rooms() {
  const [room, setRooms] = useState<RoomInterface[]>([]);

  useEffect(() => {
    Rooms();
  }, []);

  const Rooms = async () => {
    let res = await GetRoom(); console.log(res)
    if (res) {
      setRooms(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },
    {
      field: "Room_type",
      headerName: "ประเภท",
      width: 250,
      valueFormatter: (params) => params.value.Room_type_name,
    },
    {
      field: "Room_price",
      headerName: "ราคา",
      width: 150,
      valueFormatter: (params) => params.value.Price,
    },
    {
      field: "Set_of_furniture",
      headerName: "เซทเฟอนิเจอร์",
      width: 150,
      valueFormatter: (params) => params.value.Set_of_furniture_title,
    },
    
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลห้อง
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/Room/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={room}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Rooms;