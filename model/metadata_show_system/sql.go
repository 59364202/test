// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show_system

var sqlSelect = `SELECT id, metadata_show_system FROM public.lt_metadata_show_system WHERE deleted_at=to_timestamp(0)`
