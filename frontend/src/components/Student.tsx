import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { StudentInterface } from "../interfaces/IStudent";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { GetStudent } from "../services/HttpClientService";

function Student() {
  const [students, setStudent] = useState<StudentInterface[]>([]);

  const getStudent = async () => {
    let res = await GetStudent();
    if (res) { console.log(res)
      setStudent(res);
    }
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },
    { field: "STUDENT_NAME", headerName: "ชื่อ-สกุล", width: 220,  valueFormatter: (params) => params.value.STUDENT_NAME,},
    { field: "STUDENT_NUMBER", headerName: "รหัสประจำตัวนักเรียน", width: 220 , valueFormatter: (params) => params.value.STUDENT_NUMBER,},  
    { field: "Program", headerName: "แผนการเรียน", width: 100 , valueFormatter: (params) => params.value.Program_name,},    
    { field: "Gender", headerName: "เพศ", width: 200 , valueFormatter: (params) => params.value.GENDER_NAME,}, 
    { field: "Province", headerName: "จังหวัด", width: 200  ,valueFormatter: (params) => params.value.Name,},    
       
  ];

  useEffect(() => {
    getStudent();
  }, []);

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
              ข้อมูลนักศึกษา
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/student/create"
              variant="contained"
              color="inherit"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <div style={{ height: 800, width: "120%", marginTop: "20px" }}>
          <DataGrid
            rows={students}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={10}
            rowsPerPageOptions={[10]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Student;